import { ICalculatorStrategy } from "../calculator/calculator"

export class PrimeHandler implements ICalculatorStrategy {
    isOk(num: number): boolean{
        if (num ==1){
            return false
        }else if(num ==2){
            return true
        }
        for (let i=2 ; i<num ; i++){
            if (num%i==0){
                return false
            }
        }
        return true
    }
}