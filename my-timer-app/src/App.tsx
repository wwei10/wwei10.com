import React, { useEffect, useState } from 'react';
import './App.css';

class TimeLeft {
  days: number;
  hours: number;
  minutes: number;
  seconds: number;

  constructor(difference: number) {
    this.days = Math.floor(difference / (60 * 60 * 24));
    this.hours = Math.floor(difference / (60 * 60) % 24);
    this.minutes = Math.floor((difference / 60) % 60);
    this.seconds = Math.floor(difference % 60);
  }
}

function calculateTimeLeft(): TimeLeft {
  let year: number = new Date().getFullYear();
  let difference = (+new Date(`9/14/${year}`) - +new Date()) / 1000;
  return new TimeLeft(difference);
}

function App() {

  const [timeLeft, setTimeLeft] = useState(calculateTimeLeft());

  useEffect(() => {
    const timer = setTimeout(() => {
      setTimeLeft(calculateTimeLeft());
    }, 1000);
    // Clear timeout if the component is unmounted.
    return () => clearTimeout(timer);
  });

  var time = ("0" + timeLeft.hours).slice(-2);
  time += ":";
  time += ("0" + timeLeft.minutes).slice(-2);
  time += ":";
  time += ("0" + timeLeft.seconds).slice(-2);

  return (
    <div>
      <div className="countdown">{ time }</div>
    </div>
  );
}

export default App;
