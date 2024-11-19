// The render() function does the job of returning the jsx 
// that we wish to render in our application for this particular component.

import React, { Component } from "react";
import "./Funds.scss";
import Message from '../Message/Message';

class Funds extends Component {
  render() {
    console.log(this.props.funds);
    const messages = this.props.funds.map(msg => <Message message={msg.data} />);

    return (
      <div className='Funds'>
        <h2>Funds</h2>
        {messages}
      </div>
    );
  };
}

export default Funds;