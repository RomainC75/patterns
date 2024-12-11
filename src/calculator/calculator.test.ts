
import { describe, test, expect } from "@jest/globals"
import { Calculator } from "./calculator"
import { PrimeHandler } from "../math/prime"

const MIN = 1_000_000;
const MAX = 2_000_000;

describe('Sum function', () =>{
    test('test serie', () =>{
        const calculator = new Calculator(new PrimeHandler())
        let res = calculator.executeSerie(MIN, MAX)
        console.log("-> res : ", res)
        // expect(res.length).toBe(0)
    })

    test('test parallel', async () =>{
        const calculator = new Calculator(new PrimeHandler())
        let res = await calculator.executeConcurrent(MIN, MAX)
        console.log("-> res : ", res)
        // expect(res.length).toBe(0)
    })

})