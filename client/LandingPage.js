import React from 'react';
import { StyleSheet, View, Text, Picker, AsyncStorage } from 'react-native';
import { Header, Icon, FormLabel, Card, ButtonGroup } from 'react-native-elements';

class Book extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: this.props.actionType === 'SELL' ? 0 : 1
        }
        this.updateSaleOrShare = this.updateSaleOrShare.bind(this);
    }
    updateSaleOrShare(selectedIndex) {
        this.setState({
            selectedIndex
        })
    }
    render() {
        console.log(this.props);
        return (
            <Card title={this.props.name}>
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
            username: ''
        };
        this.updateOwnedFilter = this.updateOwnedFilter.bind(this);
        this.updateStatusFilter = this.updateStatusFilter.bind(this);
        this.updateList = this.updateList.bind(this);
    }
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            this.setState({ username: 'Ritesh' }, this.updateList);
        } catch (error) {
            alert(error);
        }
    }
    updateOwnedFilter(owner) {
        this.setState({ owner }, this.updateList);
    }
    updateStatusFilter(status) {
        this.setState({ status });
    }
    updateList() {
        const api =
            this.state.owner === 'owned'
                ? 'search_by_owner'
                : 'search_by_requester';
        console.log(
            `http://10.32.239.106:8080/entities/${api}/${this.state.username}`
        );
        fetch(
            `http://10.32.239.106:8080/entities/${api}/${this.state.username}`
        )
            .then(res => res.json())
            .then(books => {
                this.setState({ books });
            });
    }
    render() {
        return (
            <View style={styles.container}>
                <FormLabel>Filter By</FormLabel>
                <View flexDirection="row" justifyContent="space-evenly">
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
                        <Picker.Item label="Sale" value="sale" />
                        <Picker.Item label="Share" value="share" />
                        <Picker.Item label="All" value="all" />
                    </Picker>
                </View>
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
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff'
        // alignItems: 'flex-start',
        // justifyContent: 'center'
    }
});
