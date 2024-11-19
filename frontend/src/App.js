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
        funds: [...this.state.funds, msg]
        // 为啥下面这么写不行，输入了 Funds块的显示不更新
        // funds: [msg]
      }))
      console.log(this.state);
    });
  }

  sendAddMsg(event) {
    if (event.keyCode === 13) {
      var msg = {};
      msg["msgType"] = "Add";
      msg["msgData"] = event.target.value;
      sendMsg(JSON.stringify(msg));
      event.target.value = "";
    }
  }

  sendDelMsg(event) {
    if (event.keyCode === 13) {
      var msg = {};
      msg["msgType"] = "Del";
      msg["msgData"] = event.target.value;
      sendMsg(JSON.stringify(msg));
      event.target.value = "";
    }
  }

  render() {
    return (
      <div className="App">
        <Header />
        <Funds funds={this.state.funds} />
        <Input send={this.sendAddMsg} doc="添加" />
        <Input send={this.sendDelMsg} doc="删除" />
      </div>
    );
  }
}

export default App;
