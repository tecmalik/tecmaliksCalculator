import React from 'react'
import style from './calculator.module.css'
import CustomButton from '../reusable/CustomButton'
import {useState} from 'react'



const Calculator=()=>{

    const [value, setValue] = useState('0');


    const handleAddNumber = (e,number) =>{
        e.preventDefault();
        console.log("current value"+ value)
        if (value === '0'){
            console.log(value)
            setValue(number);
        }else{
             setValue(prevValue => prevValue + number);
        }
            
    };


    return(
        <div id ={style.wrapper}>
            <form className = {style.calculatorBody}>
                <div className = {style.display}>
                    <h1>{value}</h1>
                </div>
                <div className = {style.body}>
                    <div className = {style.firstRow}>
                        <CustomButton onClick = {() => handleAddNumber('1')} text ='1' />
                        <CustomButton onClick = {() => handleAddNumber('2')} text ='2'/>
                        <CustomButton text ='3'/>
                        <CustomButton text ='+'/>
                    </div>
                    <div className={style.secondRow}>
                        <CustomButton text ='4'/>
                        <CustomButton text ='5'/>
                        <CustomButton text ='6'/>
                        <CustomButton text ='-'/>
                    </div>
                    <div className={style.thirdRow}>
                        <CustomButton text ='7'/>
                        <CustomButton text ='8'/>
                        <CustomButton text ='9'/>
                        <CustomButton text ='*'/>
                    </div>
                    <div className={style.forthRow}>
                        <CustomButton text ='%'/>
                        <CustomButton text ='0'/>
                        <CustomButton text ='/'/>
                        <CustomButton text ='C'/>
                    </div>
                    <div className={style.lastRow}>
                    
                        <button className = {style.lastButton } > =</button>
                        
                    </div>


                </div>
            </form>  



        </div>
    

    )
    

};
export default Calculator;