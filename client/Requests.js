import React from 'react';
import { StyleSheet, View, Text, AsyncStorage } from 'react-native';
import { ButtonGroup, Card, Icon } from 'react-native-elements';

// import { BottomNavigation } from 'react-native-material-ui';

class Request extends React.Component {
    constructor(props) {
        super(props);
        this.onCancel = this.onCancel.bind(this);
        this.onApprove = this.onApprove.bind(this);
    }
    onCancel() {
        console.log('cancel request');
        fetch(`http://104.211.228.54/requests/${this.props.requestId}`, {
            method: 'PATCH',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                status: 'CANCELLED',
                ownerComment: ''
            })
        }).then(res => {
            console.log(res);
            this.props.onAction();
        });
    }
    onApprove() {
        console.log('cancel request');
        fetch(`http://104.211.228.54/requests/${this.props.requestId}`, {
            method: 'PATCH',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                status: 'APPROVED',
                ownerComment: ''
            })
        }).then(res => {
            console.log(res);
            this.props.onAction();
        });
    }
    render() {
        const { entity, type } = this.props;
        return (
            <Card>
                <Text>{entity.name}</Text>
                <Text>Owner - {entity.owner}</Text>

                {type !== 'as-owner' ? (
                    <View flexDirection="row" alignSelf="flex-end">
                        <Icon
                            raised
                            name="check"
                            color="#00ff00"
                            onPress={this.onApprove}
                        />

                        <Icon
                            raised
                            name="close"
                            color="#f50"
                            onPress={this.onCancel}
                        />
                    </View>
                ) : (
                    <View flexDirection="row" alignSelf="flex-end">
                        <Icon
                            raised
                            name="close"
                            color="#f50"
                            onPress={this.onCancel}
                        />
                    </View>
                )}
            </Card>
        );
    }
}

export default class Requests extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0,
            username: '',
            requests: [],
            loaded: false
        };
        this.getRequests = this.getRequests.bind(this);
        this.updateFilter = this.updateFilter.bind(this);
    }
    async componentDidMount() {
        try {
            const value = await AsyncStorage.getItem('username');
            this.setState({ username: value }, this.getRequests);
        } catch (error) {
            alert(error);
        }
    }
    getRequests() {
        const api = !this.state.selectedIndex
            ? 'search_by_owner'
            : 'search_by_requester';
        console.log(
            `http://104.211.228.54/requests/${api}/${this.state.username}`
        );
        fetch(
            `http://104.211.228.54/requests/${api}/${this.state.username}`
        )
            .then(res => {
                if (res.status === 404) {
                    return [];
                }
                return res.json();
            })
            .then(requests => {
                console.log(requests);
                this.setState({
                    requests: requests.filter(req => req.status === 'PENDING'),
                    loaded: true
                });
            });
    }
    updateFilter(selectedIndex) {
        this.setState({ selectedIndex }, this.getRequests);
    }
    render() {
        return (
            <View style={styles.container}>
                <ButtonGroup
                    selectedIndex={this.state.selectedIndex}
                    onPress={this.updateFilter}
                    buttons={['Incoming', 'My Open']}
                    containerStyle={{ height: 50, width: 200, marginTop: 20 }}
                />
                {this.state.loaded &&
                    !this.state.requests.length && <Text>No requests</Text>}
                {this.state.requests.map(request => (
                    <Request
                        key={request.requestId}
                        {...request}
                        type={
                            this.state.selectedIndex
                                ? 'as-owner'
                                : 'as-requester'
                        }
                        onAction={this.getRequests}
                    />
                ))}
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'flex-start'
    }
});
