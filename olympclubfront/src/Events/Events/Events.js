import React from "react";
import ListOfOlympiads from "../ListOfEvents/ListOfEvents";
import olympiad from "../Event/Event";
import * as events from "events";
import ListOfEvents from "../ListOfEvents/ListOfEvents";

class Events extends React.Component {
    constructor(props) {
        super(props);
        this.state = {events: [], token: this.props.token};
    }

    getEvents(){
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log(process.env.REACT_APP_API_URL+"/events")
        fetch(process.env.REACT_APP_API_URL+"/events", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json => {
                let events = json.events
                events.map((item, index) => {
                    fetch(process.env.REACT_APP_API_URL + "/holder/" + item.holder_id, {
                        mode: "cors",
                        credentials: "omit",
                        method: 'GET',
                        headers: headers
                    }).then(response => response.json())
                        .then(json => {
                            console.log(json)
                            let events = this.state.events
                            events[index].holder = json.holder
                            this.setState({events: events})
                        })
                        .catch(error => console.log('Authorization failed: ' + error.message))
                })
                console.log(events)
                this.setState({events: json.events})
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    componentDidMount() {
        this.getEvents()
    }

    render() {
        const containerStyle = {
            paddingTop: 120,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
        }
        return(
            <div className="container" style={containerStyle}>
                <ListOfEvents events={this.state.events} token={this.state.token}/>
            </div>
        )
    }
}

export default Events;