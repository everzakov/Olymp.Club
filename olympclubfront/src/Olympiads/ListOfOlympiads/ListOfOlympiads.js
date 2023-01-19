
import classes from './ListOfOlympiads.scss'
import {Link} from "react-router-dom";

const ListOfOlympiads = ({olympiads, token}) => {
    return (
        <div className={`${classes.olympiadContainer} olympiadContainer`}>
            {olympiads.map((item, index) => (
                <table className={`${classes.olympiad} olympiad`}>
                    <tr>
                        <td>
                            <Link className="BigOlympiadLink"
                               to={"/olympiad/" + (item.big_olympiad_id || "")}>{(item.bigOlympiad || {}).name}</Link>
                        </td>
                        <td>
                            <Link className="OlympiadLink" to={"/olympiad/" + item.big_olympiad_id + "/" + item.id}>{item.name}</Link>
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
                            <Link className="ButtonLink" to={"/olympiad/" + item.big_olympiad_id + "/" + item.id}>Подробнее</Link>
                        </td>
                    </tr>
                </table>
            ))}
        </div>
    );
}

export default ListOfOlympiads;