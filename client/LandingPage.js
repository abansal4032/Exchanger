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
    ButtonGroup
} from 'react-native-elements';

class Book extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: this.props.actionType === 'SELL' ? 0 : 1
        };
        this.updateSaleOrShare = this.updateSaleOrShare.bind(this);
    }
    updateSaleOrShare(selectedIndex) {
        this.setState(
            {
                selectedIndex
            },
            () => {
                fetch(
                    `http://10.32.239.106:8080/entities/${
                        this.props.entityId
                    }/action/${!this.props.actionType ? 'SELL' : 'SHARE'}`,
                    {
                        method: 'PATCH'
                    }
                ).then(res => console.log(res));
            }
        );
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
                <ButtonGroup
                    selectedIndex={this.state.selectedIndex}
                    onPress={this.updateSaleOrShare}
                    buttons={['For Sale', 'For Share']}
                    containerStyle={{ height: 30, width: 200 }}
                />
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
        } catch (error) {
            alert(error);
        }
    }
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
            `http://10.32.239.106:8080/entities/${api}/${
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
                    <Book key={book.name} {...book} />
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
                    onPress={() => this.props.navigation.navigate('addBook')}
                />
            </ScrollView>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        backgroundColor: '#fff',
        alignItems: 'stretch'
        // justifyContent: 'center'
    }
});
