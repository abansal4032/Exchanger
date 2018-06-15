import React from 'react';
import { View, Text, AsyncStorage } from 'react-native';
import { createStackNavigator, createSwitchNavigator } from 'react-navigation';
import NewUserForm from './NewUserForm';
import AddBook from './AddBook';
import ScanBook from './ScanBook';
import LandingPage from './LandingPage';
import BookSearchForm from './BookSearchForm';

import { Header } from 'react-native-elements';

const AppStack = createStackNavigator({
    landingPage: {
        screen: LandingPage,
        navigationOptions: () => ({
            header: (
                <Header
                    centerComponent={{
                        text: 'Pustakalaya',
                        style: { color: '#fff' }
                    }}
                />
            )
        })
    },
    addBook: {
        screen: AddBook,
        navigationOptions: () => ({
            title: 'Add New Book'
        })
    },
    scanBook: {
        screen: ScanBook,
        navigationOptions: () => ({
            title: 'Scan Book Barcode'
        })
    },
    searchBook: {
        screen: BookSearchForm,
        navigationOptions: () => ({
            header: (
                <Header
                    centerComponent={{
                        text: 'Explore',
                        style: { color: '#fff' }
                    }}
                />
            )
        })
    }
});

class AuthLoadingScreen extends React.Component {
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            this.props.navigation.navigate(value ? 'App' : 'Auth', {
                username: value
            });
        } catch (error) {
            alert(error);
        }
    }
    render() {
        return <Text>{''}</Text>;
    }
}

const App = createSwitchNavigator(
    {
        AuthLoading: AuthLoadingScreen,
        App: AppStack,
        Auth: BookSearchForm //THIS IS A HACK - Remove it
    },
    {
        initialRouteName: 'AuthLoading'
    }
);

export default App;
