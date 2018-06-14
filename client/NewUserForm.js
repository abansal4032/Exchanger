import React from 'react';
import {
    StyleSheet,
    Text,
    View,
    ToolbarAndroid,
    // FormInput,
    // Button,
    Alert
} from 'react-native';
import { FormLabel, FormInput, FormValidationMessage, Button } from 'react-native-elements'
// import { BottomNavigation } from 'react-native-material-ui';

export default class NewUserForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            alias: '',
            contactNo: '',
            points: '',
            location: ''
        };
        this.onSubmit = this.onSubmit.bind(this);
    }
    onSubmit() {
        Alert.alert(
            'Confirm Details',
            JSON.stringify(this.state),
            [
                {
                    text: 'Cancel',
                    onPress: () => console.log('Cancel Pressed'),
                    style: 'cancel'
                },
                {
                    text: 'OK',
                    onPress: () => {
                        fetch('http://10.32.239.106:8080/users', {
                            method: 'POST',
                            headers: {
                                Accept: 'application/json',
                                'Content-Type': 'application/json'
                            },
                            body: JSON.stringify({
                                name: this.state.name,
                                contact: this.state.contactNo,
                                email: 'test@gmail.com',
                                location: 1,
                                credits: 100
                            })
                        })
                        .then(res => res.json());
                    }
                }
            ],
            { cancelable: false }
        );
    }
    render() {
        return (
            <View style={styles.container}>
                <FormLabel>Name</FormLabel>
                <FormInput
                    style={{ height: 40 }}
                    onChangeText={name => this.setState({ name })}
                    value={this.state.text}
                />
                <FormLabel>Alias</FormLabel>
                <FormInput
                    style={{ height: 40 }}
                    onChangeText={alias => this.setState({ alias })}
                    value={this.state.alias}
                />
                <FormLabel>Contact Number</FormLabel>
                <FormInput
                    style={{ height: 40 }}
                    onChangeText={contactNo => this.setState({ contactNo })}
                    value={this.state.contactNo}
                />
                <FormLabel>Points</FormLabel>
                <FormInput
                    style={{ height: 40 }}
                    onChangeText={points => this.setState({ points })}
                    value={this.state.points}
                />
                <FormLabel>Location</FormLabel>
                <FormInput
                    style={{ height: 40 }}
                    onChangeText={location => this.setState({ location })}
                    value={this.state.location}
                />
                <View
                    flexDirection="row"
                    justifyContent="center"
                    style={{ width: '100%' }}
                >
                    <Button title="Submit" onPress={this.onSubmit} />
                </View>
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        // alignItems: 'flex-start',
        justifyContent: 'center'
    }
});
