import http from 'k6/http';
import { check, sleep } from 'k6';

let carId = 0;
const host = '172.23.0.3';
const port = '8080'; 

export default function () {
    const params = {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Basic ZW50cnVzdDpaVzUwY25WemREcHdZWE56TWpBeU1RPT0='
        },
      };

    createCar(params);
    getCar(params);
    deletCar(params);
    getCarNotfound(params);
    unauthorized();
    createCarBadReq(params);
}

function createCar(params){
    const url = 'http://'+ host + ':' + port + '/cars';
    const payload = JSON.stringify({
        brand: 'bmw',
        model: '431d',
        horse_power: 170
    });  
    let res = http.post(url, payload, params);
    const car = JSON.parse(res.body);
    carId = car.id;
    console.log(res.status)
    check(res, {
        'is status 200 (create car)': (r) => r.status === 200,
    });
}

function getCar(params) {
   
    const url = 'http://'+ host + ':' + port + '/cars/' + carId;
    let res = http.get(url, params);
    console.log(res.body)
    check(res, {
        'is status 200 (get car)': (r) => r.status === 200,
    });
}

function deletCar(params) {
   
    const url = 'http://'+ host + ':' + port + '/cars/' + carId;
    let res = http.del(url, null,params);
    console.log(res.body)
    check(res, {
        'is status 203 (delete car)': (r) => r.status === 203,
    });
}

function getCarNotfound(params) {
   
    const url = 'http://'+ host + ':' + port + '/cars/' + carId;
    let res = http.get(url, params);
    check(res, {
        'is status 404 (get car Not found)': (r) => r.status === 404,
    });
}

function createCarBadReq(params){
    const url = 'http://'+ host + ':' + port + '/cars';
    const payload = JSON.stringify({
        brand: 'bmw',
    });  
    let res = http.post(url, payload, params);
    const car = JSON.parse(res.body);
    check(res, {
        'is status 400 (create car bad)': (r) => r.status === 400,
    });
}unauthorized

function unauthorized() {
    const url = 'http://'+ host + ':' + port + '/cars/' + carId;
    let res = http.get(url);
    check(res, {
        'is status 401 (get car Not found)': (r) => r.status === 401,
    });
}