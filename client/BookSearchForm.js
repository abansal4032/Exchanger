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
    FormLabel,
    FormInput,
    Header,
    Icon,
    Card,
    ButtonGroup,
    Button
} from 'react-native-elements';

class Book extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0
        };
        this.updateBuyOrBorrow = this.updateBuyOrBorrow.bind(this);
    }

    updateBuyOrBorrow(selectedIndex) {
        alert('asking for ' + (selectedIndex ? 'rent' : 'buy'));
        this.setState(
            {
                selectedIndex
            },
            () => {
                fetch(`http://104.211.228.54/requests`, {
                    method: 'POST',
                    headers: {
                        Accept: 'application/json',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        entityId: this.props.entityId,
                        requester: this.props.username,
                        intent: this.state.selectedIndex ? 'RENT' : 'BUY',
                        durationInDays: 10,
                        status: 'PENDING',
                        requesterComment: 'testComment'
                    })
                }).then(res => console.log(res));
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
                    onPress={this.updateBuyOrBorrow}
                    buttons={['BUY', 'RENT']}
                    containerStyle={{ height: 30, width: 200 }}
                />
                <Text>Status: {this.props.status}</Text>
            </Card>
        );
    }
}

export default class BookSearchForm extends React.Component {
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
            fetch(`http://104.211.228.54/entities`)
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
        const searchID = this.state.name;

        fetch(`http://104.211.228.54/entities/search_by_name/${searchID}`)
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
                {/* <FormLabel>Filter By</FormLabel>*/}
                <View flexDirection="row" justifyContent="flex-start">
                    {/*<Picker
                        selectedValue={this.state.owner}
                        style={{ height: 50, width: '45%' }}
                        onValueChange={(itemValue, itemIndex) =>
                            this.updateOwnedFilter(itemValue)
                        }
                    >
                        <Picker.Item label="Borrowed" value="borrowed" />
                        <Picker.Item label="Owned" value="owned" />
                        // <Picker.Item label="All" value="all" /> 
                    </Picker> */}

                    <FormLabel>Type Book Name</FormLabel>
                    <FormInput
                        style={{ height: 40 }}
                        onChangeText={name => this.setState({ name })}
                        value={this.state.name}
                    />

                    {/* <Picker
                        selectedValue={this.state.status}
                        style={{ height: 50, width: '45%' }}
                        onValueChange={(itemValue, itemIndex) =>
                            this.updateStatusFilter(itemValue)
                        }
                    >
                        <Picker.Item label="For Sale" value="SELL" />
                        <Picker.Item label="For Share" value="SHARE" />
                        <Picker.Item label="All" value="all" />
                    </Picker> */}
                </View>

                <View flexDirection="row" style={{ width: '100%' }}>
                    <Button title="Explore" onPress={this.updateList} />
                </View>
                {this.state.loaded &&
                    !this.state.books.length && <Text>No books found</Text>}
                {this.state.books.map(book => (
                    <Book
                        key={book.entityId}
                        {...book}
                        username={this.state.username}
                    />
                ))}
                {/* <Icon
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
                /> */}
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
