import React from 'react';
import { View } from 'react-native';
import { createStackNavigator } from 'react-navigation';
import NewUserForm from './NewUserForm';

import {
  Text,
} from 'react-native';

const Routes = createStackNavigator({
    registration: {
        screen: NewUserForm,
        navigationOptions: () => ({
            title: 'Register'
        })
    }
});

export default Routes;

