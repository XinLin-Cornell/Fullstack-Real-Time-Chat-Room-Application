import React, { Component } from 'react';
import "./Message.scss";
import axios from "axios";
import "../Auth/Auth";
import { baseUrl, consoleLogger } from "../../api";

import upvoteImage from '../../assets/upvote.png';
import downvoteImage from '../../assets/downvote.png';

class Message extends Component {
  constructor(props) {
    super(props);
    const { chatRoomName, chatRoomId, chatMessage, chatUserEmail, chatUserName, Upvotes, Downvotes, MessageID } =
      this.props.message;
    this.state = {
      chatMessage,
      chatRoomId,
      chatRoomName,
      chatUserEmail, 
      chatUserName,
      Upvotes,
      Downvotes,
      MessageID,
    };
  }

  upvote = async () => {

    try {
      const response = await axios.put(`${baseUrl}/v1/api/chat/messages/${this.state.MessageID}/upvote`, {}, {
        headers: { Authorization: `Bearer ${this.props.jwtToken}`},
      });
      this.setState({ Upvotes: response.data.Upvotes });
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  downvote = async () => {
    try {
      const response = await axios.put(`${baseUrl}/v1/api/chat/messages/${this.state.MessageID}/downvote`, {}, {
        headers: { Authorization: `Bearer ${this.props.jwtToken}`},
      });
      this.setState({ Downvotes: response.data.Downvotes });
    } catch (error) {
      this.handleAxiosError(error);
    }
  }

  handleAxiosError(error) {
    consoleLogger(error);
    if (error.response) {
      if (error.response.status === 401) {
        this.props.logout();
        alert("Session expired");
        return;
      }
      this.setState({
        apiError: true,
        apiErrorMessage: error.response.data.Message || " There was an issue. Try again later.",
        submitting: false,
      });
    } else if (error.request) {
      this.setState({
        submitting: false,
        apiError: true,
        apiErrorMessage: "The request was made but no response was received",
      });
    } else {
      this.setState({
        submitting: false,
        apiError: true,
        apiErrorMessage: "Something went wrong",
      });
    }
  }

  render() {
    return (
      
      <div className="Message">
        <span className="text-decoration-underline">{this.state.chatUserName}</span>
        : {this.state.chatMessage} 
        <button onClick={this.upvote}>
        <img src={upvoteImage} alt="Upvote" className="vote-icon" />
        </button>
        {this.state.Upvotes}

        <button onClick={this.downvote}>
          <img src={downvoteImage} alt="Downvote" className="vote-icon" />
        </button>
        {this.state.Downvotes}
      </div>
    );
  }
}

export default Message;
