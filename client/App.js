import React from 'react';
import { Permissions, Notifications } from 'expo';

import {
  Text,
  View,
} from 'react-native';

// This refers to the function defined earlier in this guide

import Routes from './Routes';
const PUSH_ENDPOINT = 'http://requestbin.fullcontact.com/pg17hmpg';

export default class App extends React.Component {
  state = {
    notification: {},
  };

  componentDidMount() {
    registerForPushNotifications1();

    // Handle notifications that are received or selected while the app
    // is open. If the app was closed and then opened by tapping the
    // notification (rather than just tapping the app icon to open it),
    // this function will fire on the next tick after the app starts
    // with the notification data.
    console.log("registerForPushNotifications1")
    this._notificationSubscription = Notifications.addListener(this._handleNotification);

    console.log("addListener")
  }

  _handleNotification = (notification) => {
    this.setState({notification: notification});
    console.log("notification", notification)
  };

  render() {
    return (
        <Routes />
    );
  }
}

async function registerForPushNotifications1() {
    const { status: existingStatus } = await Permissions.getAsync(
      Permissions.NOTIFICATIONS
    );
    let finalStatus = existingStatus;
  
    // only ask if permissions have not already been determined, because
    // iOS won't necessarily prompt the user a second time.
    if (existingStatus !== 'granted') {
      // Android remote notification permissions are granted during the app
      // install, so this will only ask on iOS
      const { status } = await Permissions.askAsync(Permissions.NOTIFICATIONS);
      finalStatus = status;
    }
  
    // Stop here if the user did not grant permissions
    if (finalStatus !== 'granted') {
      return;
    }
  
    // Get the token that uniquely identifies this device
    let token = await Notifications.getExpoPushTokenAsync();
    console.log(token)
    // POST the token to your backend server from where you can retrieve it to send push notifications.
    return fetch(PUSH_ENDPOINT, {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        token: {
          value: token,
        },
        user: {
          username: 'Brent',
        },
      }),
    });
  }