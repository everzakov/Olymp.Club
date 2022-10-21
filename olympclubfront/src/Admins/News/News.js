import {Link, useNavigate} from "react-router-dom";
import classes from "./News.scss";
import {useState} from "react";
import md5 from "blueimp-md5";
import {toast, ToastContainer} from "react-toastify";
import Cookies from "universal-cookie";
import React, {useEffect} from "react";

const AddNews = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();

    const [olympiads, setOlympiads] = useState([])
    const [events, setEvents] = useState([])
    const [isOlympiad, setIsOlympiad] = useState(true)
    const [bigOlympiads, setBigOlympiads] = useState([])
    const [tableVar, setTable] = useState("Олимпиады")
    const [keyStruct, setKeyStruct] = useState(-1)


    const getBigOlympiads = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/big_olympiads")
        fetch(process.env.REACT_APP_API_URL+"/big_olympiads", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString()
            }
        })
            .then(response => response.json())
            .then(json =>{
                // console.log(json.olympiads)
                setBigOlympiads(json.olympiads)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const changeKey = (e) => {
        setKeyStruct(e[0].value)
    }

    const save = (e) => {
        // console.log("holder", e.target.holder)
        e.preventDefault()
        const formData = new FormData();
        formData.append("title", e.target.title.value);
        formData.append("description", e.target.description.value);
        formData.append("table", tableVar);
        formData.append("key", keyStruct.toString())
        // formData.append("holder-name", name);

        let headers = new Headers();
        // headers.append('Type', "formData");
        headers.append('Origin',window.location.origin.toString());
        // headers.append("Content-Type", "multipart/form-data")
        /*headers.append("Content-Type", "multipart/form-data; boundary=AaB03x" +
            "--AaB03x" +
            "Content-Disposition: file" +
            "Content-Type: png" +
            "Content-Transfer-Encoding: binary" +
            "...data... " +
            "--AaB03x--")*/
        headers.append("Authorization", "Bearer " + token)
        // headers.append("X-CSRF-Token", csrfToken)
        // console.log(process.env.REACT_APP_API_URL+"/admin/news")
        fetch(process.env.REACT_APP_API_URL+"/admin/news", {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body:formData
        })
            .then(response => {
                // console.log(response.status)
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    toast("News is created")
                }
            })
            .catch(error => {
                console.log('Authorization failed: ' + error.message)
            })


    }


    const getEvents = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/events")
        fetch(process.env.REACT_APP_API_URL+"/events", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString()
            }
        })
            .then(response => response.json())
            .then(json =>{
                // console.log(json.events)
                setEvents(json.events)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }


    const getOlympiads = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/olympiads")
        fetch(process.env.REACT_APP_API_URL+"/olympiads", {
            mode: 'cors',
            credentials: "omit",
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'Accept': 'application/json',
                'Origin': window.location.origin.toString()
            }
        })
            .then(response => response.json())
            .then(json =>{
                // console.log(json.olympiads)
                setOlympiads(json.olympiads)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getBigOlympiads()
        getEvents()
        getOlympiads()
    }, [])


    const onSiteChanged = (e) => {
        setTable(e.target.value)
        setKeyStruct(-1)
        setIsOlympiad(e.target.value == "Olympiads")
    }

    return isOlympiad ? (
        <div>
            <div className="news-window-container">
                <div className="news-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"} onSubmit={save}>
                        <p className={"news-title"}>Добавление новости</p>
                        <div className="text-input">
                            <label htmlFor="title">Заголовок новости</label>
                            <input type={"text"} name={"title"} id={"title"}/>
                        </div>
                        <div className="text-input">
                            <div>
                            <p>Олимпиады</p>
                            <input
                                type="radio"
                                name="site_name"
                                value="Olympiads"
                                onChange={onSiteChanged}
                            />
                            </div>
                            <div>
                            <p>Мероприятия</p>
                            <input
                                type="radio"
                                name="site_name"
                                value="Events"
                                onChange={onSiteChanged}
                            />
                            </div>
                        </div>
                        <div className="text-input">
                            <label htmlFor="bigOlympiad">Олимпиада</label>
                            <select name="bigOlympiad" defaultValue={olympiads.length > 0 ? olympiads[0].id: ""} multiple={false} onChange={(e) => {changeKey(e.target.selectedOptions)}}>
                                {olympiads.map((item, _) => {
                                    console.log(bigOlympiads.filter(el => {
                                        return el.id == item.big_olympiad_id
                                    })[0].name)
                                    return (
                                    <option value={item.id}>{bigOlympiads.filter(el => {
                                        return el.id == item.big_olympiad_id
                                    })[0].name} {item.name}</option>
                                )})}
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="short">Описание новости</label>
                            <input type={"text"} name={"description"} id={"description"}/>
                        </div>
                        <div className="submit-input">
                            <input id="submit" type="submit"/>
                        </div>
                    </form>
                </div>
            </div>

        </div>
    ):(
        <div>
            <div className="news-window-container">
                <div className="news-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"}>
                        <p className={"news-title"}>Добавление новости</p>
                        <div className="text-input">
                            <label htmlFor="title">Заголовок новости</label>
                            <input type={"text"} name={"title"} id={"title"}/>
                        </div>
                        <div className="text-input">
                            <div>
                                <p>Олимпиады</p>
                                <input
                                    type="radio"
                                    name="site_name"
                                    value="Olympiads"
                                    onChange={onSiteChanged}
                                />
                            </div>
                            <div>
                                <p>Мероприятия</p>
                                <input
                                    type="radio"
                                    name="site_name"
                                    value="Events"
                                    onChange={onSiteChanged}
                                />
                            </div>
                        </div>
                        <div className="text-input">
                            <label htmlFor="bigOlympiad">Мероприятие</label>
                            <select name="bigOlympiad" defaultValue={olympiads.length > 0 ? olympiads[0].id: ""} multiple={false} onChange={(e) => {changeKey(e.target.selectedOptions)}}>
                                {events.map((item, _) => {
                                    return (
                                        <option value={item.id}>{item.name}</option>
                                    )})}
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="short">Описание новости</label>
                            <input type={"text"} name={"description"} id={"description"}/>
                        </div>
                        <div className="submit-input">
                            <input id="submit" type="submit"/>
                        </div>
                    </form>
                </div>
            </div>

        </div>
    )
}

export default AddNews;