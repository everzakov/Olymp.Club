import {Link, useNavigate} from "react-router-dom";
import classes from "./Olympiad.scss";
import {useState} from "react";
import md5 from "blueimp-md5";
import {toast, ToastContainer} from "react-toastify";
import Cookies from "universal-cookie";
import React, {useEffect} from "react";

const AddOlympiad = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();
    const [image, setImage] = useState()
    const [name, setName] = useState()
    const [short, setShort] = useState()
    const [status, setStatus] = useState()
    const [olympiads, setOlympiads] = useState([])
    const [bigOlympiad, setBigOlympiad] = useState()
    const [subject, setSubject] = useState()
    const [website, setWebSite] = useState()
    const [holders, setHolders] = useState([])
    const [holder, setHolder] = useState()
const [level, setLevel] = useState()
    const [grade, setGrade] = useState("")

    let keyMap = new Map();
    keyMap.set("info", "Информатика")
    keyMap.set("math", "Математика")
    keyMap.set("econom", "Экономика")

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

    const saveStatus = (e) => {
        setStatus(e.target.value)
    }

    const saveBigOlympiad = (e) => {
        setBigOlympiad(e[0].value)
    }

    const saveHolder= (e) => {
        setHolder(e[0].value)
    }
    const saveWebSite = (e) => {
        setWebSite(e.target.value)
    }

    const saveSubject= (e) => {
        setSubject(e[0].value)
    }
    const saveLevel= (e) => {
        setLevel(e[0].value)
    }
    const saveGrade= (e) => {
        let grades = ""
        for (var i = 0; i < e.length; i++) {
            if( grades == "" ){
                grades = e[i].value
            }else {
                grades = grades + ";" + e[i].value
            }
        }
        setGrade(grades)
    }
    const save = (e) => {
       console.log("holder", e.target.holder)
        e.preventDefault()
        const formData = new FormData();
        formData.append("olympiad_logo", image);
        formData.append("name", e.target.olympiad_name.value);
        formData.append("subject", e.target.subject.value);
        formData.append("status", e.target.status.value);
        formData.append("big_olympiad", bigOlympiad);
        formData.append("short", e.target.short.value);
        formData.append("website", e.target.website.value);
        formData.append("holder", e.target.holder.value);
        formData.append("level", e.target.level.value);
        formData.append("grade", grade);

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
        console.log(process.env.REACT_APP_API_URL+"/admin/olympiad")
        fetch(process.env.REACT_APP_API_URL+"/admin/olympiad", {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body:formData
        })
            .then(response => {
                console.log(response.status)
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    toast("Olympiad is created")
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
        console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        console.log(process.env.REACT_APP_API_URL+"/holders")
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
                console.log(json.holders)
                setHolders(json.holders)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    const getOlympiads = () => {
        let headers = new Headers();
        headers.append('Content-Type', 'application/json');
        headers.append('Accept', 'application/json');
        headers.append('Origin',window.location.origin.toString());
        console.log("Bearer " + token)
        headers.append('Authorization', "Bearer " + token)
        console.log(process.env.REACT_APP_API_URL+"/big_olympiads")
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
                console.log(json.olympiads)
                setOlympiads(json.olympiads)
            })
            .catch(error => console.log('Authorization failed: ' + error.message));
    }

    useEffect(() => {
        getHolders()
        getOlympiads()
    }, [])


    return(
        <div>
            <div className="big-olympiad-window-container">
                <div className="big-olympiad-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"} onSubmit={save}>
                        <p className={"big-olympiad-title"}>Добавление олимпиады</p>
                        <div className="text-input">
                            <label htmlFor="olympiad-name">Имя олимпиады</label>
                            <input type={"text"} name={"olympiad_name"} id={"olympiad-name"} onChange={saveName}/>
                        </div>
                        <div className="text-input">
                        <label htmlFor="subject">Предмет олимпиады</label>
                        <select name="subject" defaultValue="info" multiple={false} onChange={(e) => {saveSubject(e.target.selectedOptions)}}>
                            <option value="info">Информатика</option>
                            <option value="math">Математика</option>
                            <option value="econom">Экономика</option>
                        </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="level">Уровень олимпиады</label>
                            <select name="level" defaultValue="vsosh"  multiple={false}  onChange={(e) => {saveLevel(e.target.selectedOptions)}}>
                                <option value="vsosh">ВСОШ</option>
                                <option value="1level">1 уровень</option>
                                <option value="2level">2 уровень</option>
                                <option value="3level">3 уровень</option>
                                <option value="other">Вне перечня</option>
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="grade">Класс олимпиады</label>
                            <select name="grade" defaultValue="11"  multiple={true}  onChange={(e) => {saveGrade(e.target.selectedOptions)}}>
                                <option value="11">11 класс</option>
                                <option value="10">10 класс</option>
                                <option value="9">9 класс</option>
                                <option value="8">8 класс</option>
                                <option value="7">7 класс</option>
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="bigOlympiad">Олимпиада (большая)</label>
                            <select name="bigOlympiad" defaultValue={olympiads.length > 0 ? olympiads[0].id: ""} multiple={false} onChange={(e) => {saveBigOlympiad(e.target.selectedOptions)}}>
                                {olympiads.map((item, _) => (
                                    <option value={item.id}>{item.name}</option>
                                ))}
                            </select>
                        </div>
                        <div className="text-input">
                            <label htmlFor="holder">Организатор олимпиады</label>
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
                            <label htmlFor="short">Короткое имя олимпиады</label>
                            <input type={"text"} name={"short"} id={"short"} onChange={saveShort}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="status">Статус олимпиады</label>
                            <input type={"text"} name={"status"} id={"status"} onChange={saveStatus}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="website">Сайт олимпиады</label>
                            <input type={"text"} name={"website"} id={"website"} onChange={saveWebSite}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="big-olympiad-logo">Логотип олимпиады</label>
                            <input type={"file"} name={"olympiad_logo"} id={"olympiad-logo"} onChange={saveImage}/>
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

export default AddOlympiad;