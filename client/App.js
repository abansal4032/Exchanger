import React from 'react';
import { View } from 'react-native';
import { createStackNavigator } from 'react-navigation';
import NewUserForm from './NewUserForm';

const App = createStackNavigator({
    registration: {
        screen: NewUserForm,
        navigationOptions: () => ({
            title: 'Register'
        })
    }
});

export default App;

