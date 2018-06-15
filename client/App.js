import React from 'react';
import { View, Text, AsyncStorage } from 'react-native';
import { createStackNavigator, createSwitchNavigator } from 'react-navigation';
import NewUserForm from './NewUserForm';
import AddBook from './AddBook';
import ScanBook from './ScanBook';
import LandingPage from './LandingPage';
import Requests from './Requests';
import { Header, Icon } from 'react-native-elements';

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
    requests: {
        screen: Requests,
        navigationOptions: () => ({
            title: 'Pending Requests'
        })
    }
});

class AuthLoadingScreen extends React.Component {
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            // await AsyncStorage.setItem('username', 'Avinash11');
            this.props.navigation.navigate(value ? 'App' : 'Auth');
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
        Auth: NewUserForm
    },
    {
        initialRouteName: 'AuthLoading'
    }
);

export default App;
