import {Link, useNavigate} from "react-router-dom";
import classes from "./Holder.scss";
import {useState} from "react";
import md5 from "blueimp-md5";
import {toast, ToastContainer} from "react-toastify";
import Cookies from "universal-cookie";

const AddHolder = ({token}) => {
    const navigate = useNavigate();
    const cookies = new Cookies();
    const [image, setImage] = useState()
    const [name, setName] = useState()

    const saveImage = (e) => {
        let image_as_files = e.target.files[0];
        setImage(image_as_files)
    }

    const saveName = (e) => {
        setName(e.target.value)
    }

    const save = (e) => {
        e.preventDefault()
        const formData = new FormData();
        formData.append("holder-logo", image);
        formData.append("holder-name", name);
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
        console.log(process.env.REACT_APP_API_URL+"/admin/holder")
        fetch(process.env.REACT_APP_API_URL+"/admin/holder", {
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
                    toast("Holder is created")
                }
            })
            .catch(error => {
                console.log('Authorization failed: ' + error.message)
            })


    }
    return(
        <div>
            <div className="holder-window-container">
                <div className="holder-window">
                    <ToastContainer/>
                    <form encType={"multipart/form-data"} className={"form-container"} onSubmit={save}>
                        <p className={"holder-title"}>Добавление нового организатора</p>
                        <div className="holder-name-input">
                            <label htmlFor="holder-name">Имя организатора олимпиад / мероприятия</label>
                            <input type={"text"} name={"holder_name"} id={"holder-name"} onChange={saveName}/>
                        </div>
                        <div className="holder-logo-input">
                            <label htmlFor="holder-logo">Логотип организатора</label>
                            <input type={"file"} name={"holder_logo"} id={"holder-logo"} onChange={saveImage}/>
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

export default AddHolder;