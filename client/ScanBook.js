import React from 'react';
import {
    StyleSheet,
    View,
    Text
} from 'react-native';
import { BarCodeScanner, Permissions } from 'expo';

// import { BottomNavigation } from 'react-native-material-ui';

export default class AddBook extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            hasCameraPermission: null,
        };
    }
    async componentDidMount() {
        const { status } = await Permissions.askAsync(Permissions.CAMERA);
        this.setState({hasCameraPermission: status === 'granted'});
    }
    render() {
        const { hasCameraPermission } = this.state;
    
        if (hasCameraPermission === null) {
          return <Text>Requesting for camera permission</Text>;
        } else if (hasCameraPermission === false) {
          return <Text>No access to camera</Text>;
        } else {
          return (
            <View style={{ flex: 1 }}>
              <BarCodeScanner
                onBarCodeRead={this._handleBarCodeRead}
                style={StyleSheet.absoluteFill}
              />
            </View>
          );
        }
      }
    
      _handleBarCodeRead = ({ type, data }) => {
        alert(`Bar code with type ${type} and data ${data} has been scanned!`);
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
