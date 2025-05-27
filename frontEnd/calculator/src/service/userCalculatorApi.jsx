import { createApi,fetchBaseQuery } from "@reduxjs/toolkit/query/react";


const url = "http://localhost:9090/";
export const userCalculatorApi = createApi({
    
    baseQuery : fetchBaseQuery({
    baseUrl: '/calculate',
    }),
    tagTypes:['post'],
    endpoints:(builder)=>({
        calculate: builder.mutation({
            query:(data)=>({
                url : "signup/users",
                method : "post",
                header : {"content-type" : "application/json"},
                body : data
            })
        })
    })
    
});

export const  {useCalculateMutation} = userCalculatorApi;
