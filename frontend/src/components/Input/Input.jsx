import React, { Component } from "react";
import "./Input.scss";

class Input extends Component {
  render() {
    return (
      <div className="Input">
        <span>{this.props.doc}</span>
        <input onKeyDown={this.props.send} />
      </div>
    );
  }
}

export default Input;