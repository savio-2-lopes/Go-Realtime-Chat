import React from 'react';
import Navbar from './components/Navbar/Navbar'
import Body from './components/Body/Body'

import "./styles.css"

export class App extends React.Component {
    render() {
        return (
            <div className="_main">
                <Navbar />
                <Body />
            </div>
        );
    }
}