import {Link, useNavigate} from "react-router-dom";
import classes from "./BigOlympiad.scss";
import {useState} from "react";
import md5 from "blueimp-md5";
import {toast, ToastContainer} from "react-toastify";
import Cookies from "universal-cookie";

const AddBigOlympiad = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();
    const [image, setImage] = useState()
    const [name, setName] = useState()
    const [short, setShort] = useState()
    const [description, setDescription] = useState()
    const [status, setStatus] = useState()


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

    const saveStatus = (e) => {
        setStatus(e.target.value)
    }

    const save = (e) => {
        e.preventDefault()
        const formData = new FormData();
        formData.append("big_olympiad_logo", image);
        formData.append("big_olympiad_name", name);
        formData.append("status", status);
        formData.append("short", short);
        formData.append("description", description);


        let headers = new Headers();
        headers.append('Origin',window.location.origin.toString());
        headers.append("Authorization", "Bearer " + token)

        fetch(process.env.REACT_APP_API_URL+"/admin/big_olympiad", {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body:formData
        })
            .then(response => {
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        toast(errorJson.error)
                    });
                }else{
                    toast("BigOlympiad is created")
                }
            })
            .catch(error => {
                console.log('Authorization failed: ' + error.message)
            })


    }
    return(
        <div>
            <div className="big-olympiad-window-container">
                <div className="big-olympiad-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"} onSubmit={save}>
                        <p className={"big-olympiad-title"}>Добавление олимпиады</p>
                        <div className="text-input">
                            <label htmlFor="big-olympiad-name">Имя олимпиады</label>
                            <input type={"text"} name={"big_olympiad_name"} id={"big-olympiad-name"} onChange={saveName}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="short">Короткое имя олимпиады</label>
                            <input type={"text"} name={"short"} id={"short"} onChange={saveShort}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="description">Описание олимпиады</label>
                            <input type={"text"} name={"description"} id={"description"} onChange={saveDescription}/>
                        </div>
                        <div className="text-input">
                            <label htmlFor="status">Статус олимпиады</label>
                            <input type={"text"} name={"status"} id={"status"} onChange={saveStatus}/>
                        </div>
                        <div className="holder-logo-input">
                            <label htmlFor="big-olympiad-logo">Логотип (большой) олимпиады</label>
                            <input type={"file"} name={"big_olympiad_logo"} id={"big-olympiad-logo"} onChange={saveImage}/>
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

export default AddBigOlympiad;