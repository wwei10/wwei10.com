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
  let difference = (+new Date(`10/01/${year}`) - +new Date()) / 1000;
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


  return (
    <div>
      <h1>COUNTDOWN</h1>
      <div className="countdown-container">
        <div className="countdown-el days-c">
            <p className="big-text">{ timeLeft.days }</p>
            <span>days</span>
        </div>
        <div className="countdown-el hours-c">
            <p className="big-text">{ timeLeft.hours }</p>
            <span>hours</span>
        </div>
        <div className="countdown-el mins-c">
            <p className="big-text">{ timeLeft.minutes }</p>
            <span>mins</span>
        </div>
        <div className="countdown-el seconds-c">
            <p className="big-text">{ timeLeft.seconds }</p>
            <span>seconds</span>
        </div>
      </div>
    </div>
  );
}

export default App;
