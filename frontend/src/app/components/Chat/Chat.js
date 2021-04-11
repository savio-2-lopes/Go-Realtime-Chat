import React from "react";
import { Input } from '../Input/Input';
import { Messages } from '../Messages/Messages'
import { Status } from '../Status/Status';
import { Container } from '@material-ui/core';

import './Chat.css';
const baseURL = 'ws://localhost:3333/chat';

export class Chat extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            ws: undefined,
            username: '',
            message: '',
            messages: [],
        }
    }

    render() {
        const { ws, messages } = this.state;

        return (
            <Container>
                <div className="chat">
                    <Status status={ws ? 'Online' : 'Offline'} />
                    {
                        ws && <Messages messages={messages} />
                    }

                    <div className="content_footer">
                        <div className="send-new-message">
                            <button className="addFiles">
                                <i className="fa fa-plus"></i>
                            </button>

                            <Input
                                type="text"
                                placeholder={ws ? 'Escrever mensagem' : 'Entre no chat'}
                                onChange={value => ws ? this.setMessage(value) : this.setUsername(value)}
                                defaultValue={ws ? this.state.message : this.state.username}
                            />

                            <button className="send-msg-btn" id="sendMsgBtn" onClick={() => ws ? this.sendMessage() : this.enterChat()}>
                                <i className="fa fa-paper-plane"></i>
                            </button>
                        </div>
                    </div>
                </div>
            </Container>
        )
    }

    enterChat() {
        const { username } = this.state;
        let ws = new WebSocket(baseURL + `?username=${username}`);
        ws.onopen = (evt) => {
            console.log('Conexão Socket aberta!', { evt });
        }
        ws.onclose = (evt) => {
            console.log('Conexão Socket fechada!', { evt });
            this.setState({ ws: undefined });
        }
        ws.onmessage = (msg) => {
            console.log('Mensagem da Conexão Socket:', { msg });
            this.setMessages(msg.data);
        }
        ws.onerror = (error) => {
            console.log('Erro na Conexão Socket:', { error });
        }
        this.setState({ ws, username: '' });
    }
    sendMessage() {
        const { ws, message } = this.state;
        ws.send(message);
        this.setMessage('');
    }
    setUsername(value) {
        this.setState({ username: value });
    }
    setMessage(value) {
        this.setState({ message: value });
    }
    setMessages(value) {
        let messages = this.state.messages.concat([JSON.parse(value)]);
        this.setState({ messages });
    }
}
