import {Link, Route, Routes, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import md5 from "blueimp-md5"
import Notifications from "rc-notification"
import classes from "./Request.scss";
import axios from "axios";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const PasswordRequest = () => {
    const [error, setError] = useState("")
    const [formValues, setFormValues] = useState({
        email: "",
    })
    const navigate = useNavigate();


    useEffect(() => {}, [])



    let handleChange = (e) => {
        let name = e.target.name;
        let value = e.target.value;
        formValues[name] = value;
        setFormValues(formValues);
    }

    let save = async (e) => {
        e.preventDefault();

        if (!/(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/.test(formValues.email)) {
            toast("Wrong format of email")
            return
        }

        console.log(formValues)
        let headers = new Headers();
        headers.append('Accept', 'application/x-www-form-urlencoded');
        headers.append('Origin',window.location.origin.toString());
        headers.append("Content-Type", 'application/json');
        // headers.append("X-CSRF-Token", csrfToken)
        console.log(process.env.REACT_APP_API_URL+"/password/request/post")
        console.log(JSON.stringify({
            "email": formValues.email,
        }))
        fetch(process.env.REACT_APP_API_URL+"/password/request/post", {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body: JSON.stringify({
                "email": formValues.email,
            })
        })
            .then(response => {
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    navigate("/password/request/success")
                }
            })
            .catch(error => {
                console.log('Authorization failed: ' + error.message)
            })


    }

    return(
        <div className="authorize-container">
            <div>
                <ToastContainer style={{zIndex: 10000}} />
            </div>
            <div className="authorize-window">
                <p className="authorize-title">Смена пароля</p>
                <form className="form-container" onSubmit={save} encType="">
                    <div className="email-input">
                        <label htmlFor="email">Почта</label>
                        <input id="email" type="text" name="email" onChange={handleChange}/>
                    </div>
                    <div className="submit-input">
                        <input id="submit" type="submit"/>
                    </div>

                </form>

            </div>
        </div>
    )
}
export default PasswordRequest;