export class PrimeHandler {
    calculate(num: number): boolean{
        if(num <=1){
            return true
        }
        for (let i=2 ; i<num ; i++){
            if (num%i==0){
                return true
            }
        }
        return false
    }
}