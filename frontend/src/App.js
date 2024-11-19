import React, { Component } from "react";
import './App.css';
import { connect, sendMsg } from "./api";
import Header from './components/Header/Header';
import Funds from './components/Funds/Funds';
import Input from './components/Input/Input';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      funds: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message")
      this.setState(prevState => ({
        // funds: [...this.state.funds, msg]
        funds: [msg]
      }))
      console.log(this.state);
    });
  }

  send(event) {
    if (event.keyCode === 13) {
      sendMsg(event.target.value);
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <Funds funds={this.state.funds} />
        <Input send={this.send} />
      </div>
    );
  }
}

export default App;
