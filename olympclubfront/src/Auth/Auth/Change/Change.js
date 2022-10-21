import {Link, Route, Routes, useLocation, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import md5 from "blueimp-md5"
import Notifications from "rc-notification"
import classes from "./Change.scss";
import axios from "axios";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const PasswordChange = () => {
    const search = useLocation().search;
    const token1 = new URLSearchParams(search).get('token1');
    const token2 = new URLSearchParams(search).get('token2');
    const [error, setError] = useState("")
    const [formValues, setFormValues] = useState({
        passwordFirst: "",
        passwordSecond: "",
    })
    const navigate = useNavigate();


    useEffect(() => {}, [])



    let handleChange = (e) => {
        let name = e.target.name;
        let value = e.target.value;
        formValues[name] = value;
        setFormValues(formValues);
    }

    let handleChangeSecond = (e) => {
        handleChange(e)
    }

    let save = async (e) => {
        e.preventDefault();
        if (formValues.passwordFirst.length < 8 || formValues.passwordFirst.length > 16) {
            toast("Length of password should be from 8 to 16")
            return
        }
        if (formValues["passwordFirst"] !== formValues["passwordSecond"]) {
            toast("Passwords should be the same")
            return
        }
        // console.log(formValues)
        let headers = new Headers();
        headers.append('Accept', 'application/x-www-form-urlencoded');
        headers.append('Origin',window.location.origin.toString());
        headers.append("Content-Type", 'application/json');
        // headers.append("X-CSRF-Token", csrfToken)
        // console.log(process.env.REACT_APP_API_URL+"/password/change/post?token1=" + token1 + "&token2="+token2 )
        // console.log(JSON.stringify({
        //     "email": formValues.email,
        // }))
        fetch(process.env.REACT_APP_API_URL+"/password/change/post?token1=" + token1 + "&token2="+token2 , {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body: JSON.stringify({
                "pass_hash": md5(formValues.passwordFirst),
            })
        })
            .then(response => {
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        // console.log(errorJson.error)
                        toast(errorJson.error)
                    });
                }else{
                    navigate("/password/change/success")
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
                    <div className="password-input">
                        <label htmlFor="password">Новый Пароль</label>
                        <input id="password" type="password" name="passwordFirst" onChange={handleChange}/>
                    </div>
                    <div className="password-input">
                        <label htmlFor="password">Новый Пароль ещё раз</label>
                        <input id="password" type="password" name="passwordSecond" onChange={handleChangeSecond}/>
                    </div>
                    <div className="submit-input">
                        <input id="submit" type="submit"/>
                    </div>

                </form>

            </div>
        </div>
    )
}
export default PasswordChange;