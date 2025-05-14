import React from 'react'
import style from './calculator.module.css'
import CustomButton from '../reusable/CustomButton'
import {useState} from 'react'



const Calculator=()=>{

    const [value, setValue] = useState('0');


    const handleAddNumber = (number) =>{
        
        console.log("current value"+ value)
        if (value === '0'){
            console.log(value)
            setValue(number);
        }else{
             setValue(prevValue => prevValue + number);
        }
            
    };

    const handleClearNumber = ()=>{
        setValue('0');
    }

    const handleSubmit = (e) =>{
        e.preventDefault();
    }


    return(
        <div id ={style.wrapper}>
            <form className = {style.calculatorBody} onSubmit={handleSubmit}>
                <div className = {style.display}>
                    <h1>{value}</h1>
                </div>
                <div className = {style.body}>
                    <div className = {style.firstRow}>
                        <CustomButton onClick = {() => handleAddNumber('1')} text ='1' />
                        <CustomButton onClick = {() => handleAddNumber('2')} text ='2'/>
                        <CustomButton onClick = {() => handleAddNumber('3')}  text ='3'/>
                        <CustomButton onClick = {() => handleAddNumber('+')}  text ='+'/>
                    </div>
                    <div className={style.secondRow}>
                        <CustomButton onClick = {() => handleAddNumber('4')}  text ='4'/>
                        <CustomButton onClick = {() => handleAddNumber('5')}  text ='5'/>
                        <CustomButton onClick = {() => handleAddNumber('6')}  text ='6'/>
                        <CustomButton onClick = {() => handleAddNumber('-')}  text ='-'/>
                    </div>
                    <div className={style.thirdRow}>
                        <CustomButton onClick = {() => handleAddNumber('7')}  text ='7'/>
                        <CustomButton onClick = {() => handleAddNumber('8')}  text ='8'/>
                        <CustomButton onClick = {() => handleAddNumber('9')}  text ='9'/>
                        <CustomButton onClick = {() => handleAddNumber('*')}  text ='*'/>
                    </div>
                    <div className={style.forthRow}>
                        <CustomButton onClick = {() => handleAddNumber('%')}  text ='%'/>
                        <CustomButton onClick = {() => handleAddNumber('0')}  text ='0'/>
                        <CustomButton onClick = {() => handleAddNumber('/')}  text ='/'/>
                        <button className ={style.cancel} onClick = {() => handleClearNumber()} >c</button>
                        
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