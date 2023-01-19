
import {Link } from "react-router-dom";
import Cookies from 'universal-cookie';
import {useState} from "react"
import classes from "./header.scss";

const Header = ({token}) => {
    if (token == undefined) {
        return (
            <div className="base-bar">
                <ul>
                    <li>
                        <Link to="/">
                            <img className="base-bar-logo" src={process.env.REACT_APP_STATIC_FILES + "/icons/logo.png"} alt="Лого Olymp.Club"/>
                        </Link>
                    </li>
                    <li>
                        <Link to="/olympiads">Олимпиады</Link>
                    </li>
                    <li>
                        <Link to="/events">Мероприятия</Link>
                    </li>
                    <li>
                        <Link to="/authorize">Авторизоваться</Link>
                    </li>
                </ul>
            </div>
        )
    }else{
        return (
        <div className="base-bar">
            <ul>
                <li>
                    <Link to="/">
                        <img className="base-bar-logo" src={process.env.REACT_APP_STATIC_FILES + "/icons/logo.png"} alt="Лого Olymp.Club"/>
                    </Link>
                </li>
                <li>
                    <Link to="/olympiads">Олимпиады</Link>
                </li>
                <li>
                    <Link to="/events">Мероприятия</Link>
                </li>
                <li>
                    <Link to="/me">Личный кабинет</Link>
                </li>
            </ul>
        </div>
        )
    }
}

export default Header;