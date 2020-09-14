import React, { useEffect, useState } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  useParams
} from "react-router-dom";
import './App.css';

class TimeLeft {
  days: number;
  hours: number;
  minutes: number;
  seconds: number;
  expired: boolean;

  constructor(difference: number) {
    if (difference > 0) {
      this.expired = false;
    } else {
      this.expired = true;
    }
    this.days = Math.floor(difference / (60 * 60 * 24));
    this.hours = Math.floor(difference / (60 * 60) % 24);
    this.minutes = Math.floor((difference / 60) % 60);
    this.seconds = Math.floor(difference % 60);
  }
}

interface ParamTypes {
  endtime?: string
}

type TimerProps = { endTime?: number };

function Timer(props: TimerProps) {
  let endTimeParameter: number;
  if (props.endTime === undefined) {
    endTimeParameter = +new Date() / 1000 + 30 * 60;
  } else {
    endTimeParameter = +props.endTime;
  }
  const [endTime] = useState(endTimeParameter);
  const [timeLeft, setTimeLeft] = useState(new TimeLeft(endTime - +new Date() / 1000));

  useEffect(() => {
    const timer = setTimeout(() => {
      let difference: number = endTime - +new Date() / 1000;
      setTimeLeft(new TimeLeft(difference));
    }, 1000);
    // Clear timeout if the component is unmounted.
    return () => clearTimeout(timer);
  });

  var timeLeftString = "";
  if (timeLeft.expired) {
    timeLeftString = "Expired";
  } else {
    if (timeLeft.days > 0) {
      timeLeftString += ("0" + timeLeft.days).slice(-2);
      timeLeftString += ":";
    }
    if (timeLeft.hours > 0) {
      timeLeftString += ("0" + timeLeft.hours).slice(-2);
      timeLeftString += ":";
    }
    timeLeftString += ("0" + timeLeft.minutes).slice(-2);
    timeLeftString += ":";
    timeLeftString += ("0" + timeLeft.seconds).slice(-2);
  }

  return (
      <div>
        <div className="countdown">{ timeLeftString }</div>
      </div>
  )
}

function Child() {
  let { endtime } = useParams<ParamTypes>();
  if (endtime === undefined) {
    return (
      <Timer />
    )
  }
  return (
    <Timer endTime={ +endtime } />
  )
}

function App() {
  return (
    <Router>
      <Switch>
        <Route exact path="/" children={<Timer />} />
        <Route path="/:endtime" children={<Child />} />
      </Switch>
    </Router>
  );
}

export default App;
