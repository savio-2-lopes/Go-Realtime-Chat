import React, { Component } from "react"
import { Chat } from '../Chat/Chat'
import './Body.css'

export default class Body extends Component {
    render() {
        return (
            <div className="main_body" >
                <Chat />
            </div >
        )
    }
}