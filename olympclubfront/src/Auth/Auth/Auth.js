import {Link, Route, Routes, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import md5 from "blueimp-md5"
import classes from "./Auth.scss";
import Cookies from 'universal-cookie'
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const Auth = ({setToken}) => {
    const cookies = new Cookies();
    const [error, setError] = useState("")
    const [formValues, setFormValues] = useState({
        email: "",
        password: ""
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
        let headers = new Headers();
        headers.append('Accept', 'application/x-www-form-urlencoded');
        headers.append('Origin',window.location.origin.toString());
        headers.append("Content-Type", 'application/json');

        fetch(process.env.REACT_APP_API_URL+"/auth/post", {
            mode: 'cors',
            credentials: "omit",
            method: 'POST',
            headers: headers,
            body: JSON.stringify({
                "email": formValues.email,
                "pass_hash": md5(formValues.password),
            })
        })
            .then(response => {
                if (response.status !== 200) {
                    response.json().then((errorJson) => {
                        toast(errorJson.error)
                    });
                }else{
                    response.json().then((json) => {
                        cookies.set('token', json.token, { path: '/' });
                        setToken(json.token)
                        navigate("/me")
                    });
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
                <p className="authorize-title">Авторизация</p>
                <form className="form-container" onSubmit={save} encType="">
                    <div className="email-input">
                        <label htmlFor="email">Почта</label>
                        <input id="email" type="text" name="email" onChange={handleChange}/>
                    </div>
                    <div className="password-input">
                        <label htmlFor="password">Пароль</label>
                        <input id="password" type="password" name="password" onChange={handleChange}/>
                    </div>
                    <div className="change_password" style={{width: "500px", margin: "auto auto 10px"}}>
                        <Link to="/password/request"
                              className="bar-link-hover-underline" style={{color: "#0066FF"}}>Забыли пароль?</Link>
                    </div>
                    <div className="authorize" style={{width: "500px", margin: "auto auto 10px"}}>
                        <Link to="/register" className="bar-link-hover-underline"
                              style={{color: "#0066FF"}}>Нет аккаунта?</Link>
                    </div>
                    <div className="submit-input">
                        <input id="submit" type="submit"/>
                    </div>
                </form>

            </div>
        </div>
    )
}

export default Auth;