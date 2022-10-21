import {Link, useNavigate} from "react-router-dom";
import classes from "./Event.scss";
import {useState} from "react";
import md5 from "blueimp-md5";
import {toast, ToastContainer} from "react-toastify";
import Cookies from "universal-cookie";
import React, {useEffect} from "react";

const AddEvent = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();
    const [image, setImage] = useState()
    const [name, setName] = useState()
    const [short, setShort] = useState()
    const [description, setDescription] = useState()
    const [status, setStatus] = useState()
    const [holders, setHolders] = useState([])
    const [holder, setHolder] = useState()
const [website, setWebsite] = useState()

    const saveImage = (e) => {
        let image_as_files = e.target.files[0];
        setImage(image_as_files)
    }

    const saveName = (e) => {
        setName(e.target.value)
    }

    const saveShort = (e) => {
        setShort(e.target.value)
    }

    const saveDescription = (e) => {
        setDescription(e.target.value)
    }

    const saveHolder= (e) => {
        setHolder(e[0].value)
    }

    const saveStatus = (e) => {
        setStatus(e.target.value)
    }

    const saveWebsite =(e) => {
        setWebsite(e.target.value)
    }

    const save = (e) => {
        e.preventDefault()
        const formData = new FormData();
        formData.append("event_logo", image);
        formData.append("name", name);
        formData.append("status", status);
        formData.append("short", short);
        formData.append("description", description);
        formData.append("website", website);
        formData.append("holder", e.target.holder.value);

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
        // console.log(process.env.REACT_APP_API_URL+"/admin/event")
        fetch(process.env.REACT_APP_API_URL+"/admin/event", {
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
                        // console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    toast("Event is created")
                }
            })
            .catch(error => {
                console.log('Authorization failed: ' + error.message)
            })


    }

    const getHolders = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        // console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        // console.log(process.env.REACT_APP_API_URL+"/holders")
        fetch(process.env.REACT_APP_API_URL+"/holders", {
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
                // console.log(json.holders)
                setHolders(json.holders)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getHolders()
    }, [])

    return(
        <div>
            <div className="event-window-container">
                <div className="event-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"} onSubmit={save}>
                        <p className={"event-title"}>Добавление мероприятия</p>
                        <div className="text-input">
                            <label htmlFor="event-name">Имя мероприятия</label>
                            <input type={"text"} name={"event_name"} id={"event-name"} onChange={saveName}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="holder">Организатор мероприятия</label>
                            <select name="holder" defaultValue={holders.length > 0 ? holders[0].id: ""}   multiple={false} onChange={(e) => {saveHolder(e.target.selectedOptions)}}>
                                {holders.map((item, _) => {
                                    return (
                                        <option value={item.id}>{item.name}</option>
                                    )
                                })}
                                }
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="short">Короткое имя мероприятия</label>
                            <input type={"text"} name={"short"} id={"short"} onChange={saveShort}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="description">Описание мероприятия</label>
                            <input type={"text"} name={"description"} id={"description"} onChange={saveDescription}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="status">Статус мероприятия</label>
                            <input type={"text"} name={"status"} id={"status"} onChange={saveStatus}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="big-olympiad-logo">Сайт мероприятия</label>
                            <input type={"text"} name={"website"} id={"website"} onChange={saveWebsite}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="event-logo">Логотип мероприятия</label>
                            <input type={"file"} name={"event_logo"} id={"event-logo"} onChange={saveImage}/>
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

export default AddEvent;