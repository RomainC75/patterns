export interface ICalculatorStrategy{
    isOk(num: number): boolean
}

export class Calculator {
    strategy: ICalculatorStrategy;
    
    constructor(strategy: ICalculatorStrategy){
        this.strategy = strategy
    }

    executeSerie(min: number, max: number): number[]{
        console.log('this . strategy : ', this.strategy)
        if (!this.strategy){
            throw Error("no strategy set")
        }
        const result: number[]=[];
        for(let i=min; i<max; i++){
            if (this.strategy.isOk(i)){
                result.push(i)
            }
        }
        return result
    }

    async executeConcurrent(min: number, max: number): Promise<number[]>{
        if (!this.strategy){
            throw Error("no strategy set")
        }
        // const ln = max - min;
        // const detailNumbers: number[] = Array(ln).map((n, index)=>{
        //     return index+min;
        // })
        const numbers: number[] = [];
        for(let i=min ; i<max ; i++ ){
            numbers.push(i)
        }

        const res = await Promise.allSettled(numbers.map(num=> new Promise((resolve, reject)=>{
            if (this.strategy.isOk(num)){
                resolve(num)
            }
            reject(num)
        }))) as PromiseSettledResult<number>[]
        return res.filter(p=>p.status==="fulfilled").map(fp=>fp.value)
        
    }
}