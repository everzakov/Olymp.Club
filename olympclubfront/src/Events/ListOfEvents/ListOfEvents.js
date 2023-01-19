
import classes from './ListOfEvents.scss'
import {Link} from "react-router-dom";

const ListOfEvents = ({events, token}) => {
    return (
        <div className={`${classes.olympiadContainer} olympiadContainer`}>
            {events.map((item, index) => (
                <table className={`${classes.olympiad} olympiad`}>
                    <tr>
                        <td colSpan={2}>
                            <Link className="OlympiadLink" to={"/event/" + item.id}>{item.name}</Link>
                        </td>
                        <td>
                            <p dangerouslySetInnerHTML={{__html: item.status}}></p>
                        </td>
                    </tr>
                    <tr>
                        <td rowSpan={2}>
                            <div className={`${classes.olympiadImg} olympiadImg`}>
                                <img src={process.env.REACT_APP_STATIC_FILES + "/" + item.img}></img>
                            </div>
                        </td>
                        <td>
                        </td>
                        <td>
                        </td>
                    </tr>
                    <tr>
                        <td>
                        </td>
                        <td>
                            <Link className="ButtonLink" to={"/event/" + item.id}>Подробнее</Link>
                        </td>
                    </tr>
                </table>
            ))}
        </div>
    );
}

export default ListOfEvents;