import React from 'react';
import {
    StyleSheet,
    ScrollView,
    View,
    Text,
    Picker,
    AsyncStorage,
    Image
} from 'react-native';
import {
    Header,
    Icon,
    FormLabel,
    Card,
    ButtonGroup,
    Button
} from 'react-native-elements';
import { Permissions, Notifications } from 'expo';

const PUSH_ENDPOINT = 'http://104.211.228.54/users/';

class Book extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: this.props.actionType === 'SELL' ? 0 : 1
        };
        this.updateSaleOrShare = this.updateSaleOrShare.bind(this);
        this.releaseBook = this.releaseBook.bind(this);
    }
    updateSaleOrShare(selectedIndex) {
        this.setState(
            {
                selectedIndex
            },
            () => {
                fetch(
                    `http://104.211.228.54/entities/${
                        this.props.entityId
                    }/action/${!this.props.actionType ? 'SELL' : 'SHARE'}`,
                    {
                        method: 'PATCH'
                    }
                ).then(res => console.log(res));
            }
        );
    }
    releaseBook() {
        fetch(`http://104.211.228.54/entities/${this.props.entityId}/release`, {
            method: 'PATCH'
        }).then(this.props.updateList);
    }
    render() {
        console.log(this.props);
        return (
            <Card title={this.props.name}>
                <Image
                    style={{ height: 300, width: '100%' }}
                    source={{
                        uri: `http://covers.openlibrary.org/b/isbn/${
                            this.props.attributes.isbn
                        }-M.jpg`
                    }}
                    resizeMode="cover"
                />
                {this.props.bookStatus === 'owned' ? (
                    <ButtonGroup
                        selectedIndex={this.state.selectedIndex}
                        onPress={this.updateSaleOrShare}
                        buttons={['For Sale', 'For Share']}
                        containerStyle={{ height: 30, width: 200 }}
                    />
                ) : (
                    <Button title="Release" onPress={this.releaseBook} />
                )}
                <Text>Status: {this.props.status}</Text>
            </Card>
        );
    }
}

export default class LandingPage extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            owner: 'owned',
            status: 'all',
            books: [],
            username: '',
            loaded: false
        };
        this.updateOwnedFilter = this.updateOwnedFilter.bind(this);
        this.updateStatusFilter = this.updateStatusFilter.bind(this);
        this.updateList = this.updateList.bind(this);
    }
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            this.setState({ username: value }, this.updateList);
            registerForPushNotifications1(value);
            this._notificationSubscription = Notifications.addListener(
                this._handleNotification
            );
        } catch (error) {
            alert(error);
        }
    }
    _handleNotification = notification => {
        this.setState({ notification: notification });
        console.log('notification', notification);
    };
    updateOwnedFilter(owner) {
        this.setState({ owner }, this.updateList);
    }
    updateStatusFilter(status) {
        this.setState({ status }, this.updateList);
    }
    updateList() {
        const filterPostFix =
            this.state.status === 'all' ? '' : `?filter=${this.state.status}`;
        const api =
            this.state.owner === 'owned'
                ? 'search_by_owner'
                : 'search_by_requester';

        fetch(
            `http://104.211.228.54/entities/${api}/${
                this.state.username
            }${filterPostFix}`
        )
            .then(res => {
                if (res.status === 404) {
                    return [];
                }
                return res.json();
            })
            .then(books => {
                console.log(books);
                this.setState({ books, loaded: true });
            });
    }
    render() {
        return (
            <ScrollView contentContainerStyle={styles.container}>
                <FormLabel>Filter By</FormLabel>
                <View flexDirection="row" justifyContent="flex-start">
                    <Picker
                        selectedValue={this.state.owner}
                        style={{ height: 50, width: '45%' }}
                        onValueChange={(itemValue, itemIndex) =>
                            this.updateOwnedFilter(itemValue)
                        }
                    >
                        <Picker.Item label="Borrowed" value="borrowed" />
                        <Picker.Item label="Owned" value="owned" />
                        {/* <Picker.Item label="All" value="all" /> */}
                    </Picker>

                    <Picker
                        selectedValue={this.state.status}
                        style={{ height: 50, width: '45%' }}
                        onValueChange={(itemValue, itemIndex) =>
                            this.updateStatusFilter(itemValue)
                        }
                    >
                        <Picker.Item label="For Sale" value="SELL" />
                        <Picker.Item label="For Share" value="SHARE" />
                        <Picker.Item label="All" value="all" />
                    </Picker>
                </View>
                {this.state.loaded &&
                    !this.state.books.length && <Text>No books found</Text>}
                {this.state.books.map(book => (
                    <Book
                        key={book.name}
                        {...book}
                        bookStatus={this.state.owner}
                        updateList={this.updateList}
                    />
                ))}
                <Icon
                    containerStyle={{
                        position: 'absolute',
                        right: 20,
                        bottom: 20
                    }}
                    reverse
                    raised
                    name="add"
                    color="#f50"
                    onPress={() =>
                        this.props.navigation.navigate('addBook', {
                            updateList: this.updateList
                        })
                    }
                />
                <Icon
                    containerStyle={{
                        position: 'absolute',
                        right: 20,
                        bottom: 80
                    }}
                    reverse
                    raised
                    name="search"
                    color="#ccc"
                    onPress={() => this.props.navigation.navigate('searchBook')}
                />
                <Icon
                    containerStyle={{
                        position: 'absolute',
                        right: 20,
                        bottom: 140
                    }}
                    reverse
                    raised
                    name="beenhere"
                    color="#f50"
                    onPress={() => this.props.navigation.navigate('requests')}
                />
            </ScrollView>
        );
    }
}

async function registerForPushNotifications1(userId) {
    const { status: existingStatus } = await Permissions.getAsync(
        Permissions.NOTIFICATIONS
    );
    let finalStatus = existingStatus;

    // only ask if permissions have not already been determined, because
    // iOS won't necessarily prompt the user a second time.
    if (existingStatus !== 'granted') {
        // Android remote notification permissions are granted during the app
        // install, so this will only ask on iOS
        const { status } = await Permissions.askAsync(
            Permissions.NOTIFICATIONS
        );
        finalStatus = status;
    }

    // Stop here if the user did not grant permissions
    if (finalStatus !== 'granted') {
        return;
    }

    // Get the token that uniquely identifies this device
    let token = await Notifications.getExpoPushTokenAsync();
    console.log('&&&&&&&&&&&&&&&', token);
    // POST the token to your backend server from where you can retrieve it to send push notifications.
    return fetch(PUSH_ENDPOINT + userId + '/updateToken', {
        method: 'PATCH',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            registrationToken: token
        })
    });
}

const styles = StyleSheet.create({
    container: {
        minHeight: 500,
        backgroundColor: '#fff',
        alignItems: 'stretch'
        // justifyContent: 'center'
    }
});
