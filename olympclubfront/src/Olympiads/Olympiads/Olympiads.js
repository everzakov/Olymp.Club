import React from "react";
import ListOfOlympiads from "../ListOfOlympiads/ListOfOlympiads";
import olympiad from "../Olympiad/Olympiad";

class Olympiads extends React.Component {
    constructor(props) {
        super(props);
        this.state = {olympiads: [], filter: {subject: "", level: "", grade: ""}, token: props.token};
    }

    getFilterQuery() {
        let query = "";
        if (this.state.filter.subject !== "") {
            query = "?subject=" + this.state.filter.subject;
        }
        if (this.state.filter.level !== ""){
            if (query === "") {
                query = "?"
            }else{
                query += "&"
            }
            query += "level=" + this.state.filter.level;
        }
        if (this.state.filter.grade !== "") {
            if (query === "") {
                query = "?"
            }else{
                query += "&"
            }
            query += "grade=" + this.state.filter.grade;
        }
        console.log(query)
        return query
    }

    getOlympiads(){
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        let query = this.getFilterQuery()
        console.log(process.env.REACT_APP_API_URL+"/olympiads" + query)
        fetch(process.env.REACT_APP_API_URL+"/olympiads" + query, {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: headers
        })
            .then(response => response.json())
            .then(json => {
                let olympiads = json.olympiads
                olympiads.map((item, index) => {
                    fetch(process.env.REACT_APP_API_URL + "/olympiad/" + item.big_olympiad_id, {
                        mode: "cors",
                        credentials: "omit",
                        method: 'GET',
                        headers: headers
                    }).then(response => response.json())
                        .then(json => {
                            console.log(json)
                            let olympiads = this.state.olympiads
                            olympiads[index].bigOlympiad = json.big_olympiad
                            this.setState({olympiads: olympiads})
                        })
                        .catch(error => console.log('Authorization failed: ' + error.message))
                })
                console.log(olympiads)
                this.setState({olympiads: json.olympiads})
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    componentDidMount() {
        this.getOlympiads()
    }

    changeSubjectFilter(options) {
        let filter = this.state.filter
        filter.subject = options[0].value
        this.setState({filter: filter})
        this.getOlympiads()
    }

    changeLevelFilter(options) {
        let filter = this.state.filter
        filter.level = options[0].value
        this.setState({filter:filter})
        this.getOlympiads()
    }

    changeGradeFilter(options) {
        let filter = this.state.filter
        filter.grade = options[0].value
        this.setState({filter:filter})
        this.getOlympiads()
    }

    render() {
        let map = new Map();
        map.set("info", "Информатика")
        map.set("math", "Математика")
        map.set("econom", "Экономика")
        const containerStyle = {
            paddingTop: 120,
            display: "flex",
            flexDirection: "row",
            alignItems: "center",
            justifyContent: "center",
        }
        return(
            <div >
                <form className="container" style={containerStyle}>
                    <label className="Filter">
                        <select  multiple={false} value={this.state.subject} onChange={(e)=> {this.changeSubjectFilter(e.target.selectedOptions)}}>
                            <option value="">Предмет</option>
                            <option value="info">Информатика</option>
                            <option value="math">Математика</option>
                            <option value="econom">Экономика</option>
                        </select>
                    </label>
                    <label className="Filter">
                        <select  multiple={false} value={this.state.level} onChange={(e) => {this.changeLevelFilter(e.target.selectedOptions)}}>
                            <option value="">Уровень</option>
                            <option value="vsosh">ВСОШ</option>
                            <option value="1level">1 уровень</option>
                            <option value="2level">2 уровень</option>
                            <option value="3level">3 уровень</option>
                            <option value="other">Вне перечня</option>
                        </select>
                    </label>
                    <label className="Filter">
                        <select  multiple={false} value={this.state.grade} onChange={(e) => {this.changeGradeFilter(e.target.selectedOptions)}}>
                            <option value="">Класс</option>
                            <option value="11">11 класс</option>
                            <option value="10">10 класс</option>
                            <option value="9">9 класс</option>
                            <option value="8">8 класс</option>
                            <option value="7">7 класс</option>
                        </select>
                    </label>
                </form>
                <ListOfOlympiads olympiads={this.state.olympiads} token={this.state.token}/>
            </div>
        )
    }
}

export default Olympiads;