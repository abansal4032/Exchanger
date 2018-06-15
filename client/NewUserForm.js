import React from 'react';
import {
    StyleSheet,
    Text,
    View,
    ToolbarAndroid,
    KeyboardAvoidingView,
    AsyncStorage,
    // FormInput,
    // Button,
    Alert
} from 'react-native';
import {
    FormLabel,
    FormInput,
    FormValidationMessage,
    Header,
    Button
} from 'react-native-elements';
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
        const { navigate } = this.props.navigation;
        Alert.alert(
            'Confirm Details',
            JSON.stringify(this.state),
            [
                {
                    text: 'Cancel',
                    onPress: () => navigate('addBook'),
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
                        }).then(
                            res =>
                                res.status === 200 &&
                                AsyncStorage.setItem(
                                    'username',
                                    this.state.name
                                )
                        );
                    }
                }
            ],
            { cancelable: false }
        );
    }
    render() {
        return (
            <KeyboardAvoidingView style={styles.container}>
                <Header
                    centerComponent={{
                        text: 'Register',
                        style: { color: '#fff' }
                    }}
                />
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
                    keyboardType="phone-pad"
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
                    // flexDirection="row"
                    // justifyContent="center"
                    style={{ width: '100%' }}
                >
                    <Button icon={{name: 'send'}} title="Submit" onPress={this.onSubmit} />
                </View>
            </KeyboardAvoidingView>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        paddingTop: '20px',
        flex: 1,
        backgroundColor: '#fff'
        // alignItems: 'flex-start',
        // justifyContent: 'center'
    }
});
