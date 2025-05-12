import React from 'react'
import style from './calculator.module.css'
import CustomButton from '../reusable/CustomButton'



const Calculator = () =>{
    const Value = 0

    return(
        <div id ={style.wrapper}>
            <form className = {style.calculatorBody}>
                <div className = {style.display}>
                    <h1>0</h1>
                </div>
                <div className = {style.body}>
                    <div className = {style.firstRow}>
                        <CustomButton text ='1'/>
                        <CustomButton text ='2'/>
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
                        {/* <CustomButton text ='='/> */}
                        <button className = {style.lastButton}>=</button>
                        
                    </div>


                </div>
            </form>  



        </div>
    

    )
    

}
export default Calculator