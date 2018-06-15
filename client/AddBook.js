import React from 'react';
import { StyleSheet, View, Text, Image, AsyncStorage } from 'react-native';
import {
    FormLabel,
    FormInput,
    FormValidationMessage,
    Button,
    Card
} from 'react-native-elements';
import { BarCodeScanner, Permissions } from 'expo';

// import { BottomNavigation } from 'react-native-material-ui';

export default class AddBook extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            isbn: '',
            book: null,
            username: ''
        };
        this.setIsbn = this.setIsbn.bind(this);
        this.getBookDetails = this.getBookDetails.bind(this);
        this.addBook = this.addBook.bind(this);
    }
    setIsbn(isbn) {
        this.setState({ isbn });
    }
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            this.setState({ username: value });
        } catch (error) {
            alert(error);
        }
    }
    getBookDetails() {
        console.log(this.state.isbn);
        fetch(
            `https://openlibrary.org/api/books?bibkeys=ISBN:` +
                this.state.isbn +
                `&jscmd=data&format=json`
        )
            .then(res => res.json())
            .then(data => {
                console.log(data);
                return data[`ISBN:${this.state.isbn}`];
            })
            .then(book =>
                this.setState({
                    book: {
                        title: book.title,
                        cover: book.cover && book.cover.medium,
                        by: book.by_statement
                    }
                })
            );
    }
    addBook() {
        fetch('http://10.32.239.106:8080/entities', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: this.state.book.title,
                type: 'book',
                owner: this.state.username,
                actionType: 'SELL',
                status: 'Available',
                price: 50,
                location: 1,
                attributes: {
                    isbn: this.state.isbn
                }
            })
        }).then(res => {
            console.log(res);
            this.props.navigation.goBack();
        });
    }
    render() {
        const { navigate } = this.props.navigation;
        return (
            <View style={styles.container} paddingTop={20}>
                <Button
                    title="Scan Barcode"
                    onPress={() =>
                        navigate('scanBook', { setIsbn: this.setIsbn })
                    }
                />
                <FormLabel>OR</FormLabel>
                <FormLabel>Enter ISBN below</FormLabel>
                <View
                    display="flex"
                    flexDirection="row"
                    justifyContent="space-evenly"
                >
                    <FormInput
                        value={this.state.isbn}
                        keyboardType="numeric"
                        onChangeText={isbn => this.setState({ isbn })}
                        containerStyle={{ width: '80%' }}
                    />
                    <Button
                        icon={{ name: 'send' }}
                        onPress={this.getBookDetails}
                    />
                </View>
                {this.state.book && (
                    <Card title={this.state.book.title}>
                        <Image
                            style={{ height: 300, width: '100%' }}
                            source={{ uri: this.state.book.cover }}
                            resizeMode="cover"
                        />
                        <Text>{this.state.book.by}</Text>
                        <Button title="Add this book" onPress={this.addBook} />
                    </Card>
                )}
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff'
        // alignItems: 'center',
        // justifyContent: 'center'
    }
});
