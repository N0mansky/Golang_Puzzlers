package main

import "fmt"

func main(){

    // 示例 1
    {
        type MyString = string 
        str := "BCD"
        myStr1 := MyString(str)
        myStr2 := MyString("A"+str)
        fmt.Printf("%T(%q) == %T(%q): %v\n",str,str,myStr1,myStr1,str == myStr1)

        fmt.Printf("%T(%q) > %T(%q): %v\n",str,str,myStr2,myStr2,str > myStr2)

        fmt.Printf("Type %T is the same as type %T.\n",myStr1,str)


        strs := []string{"E","F","G"}
        myStrs :=  []MyString(strs)
        fmt.Printf("A value of type []MyString: %T(%q)\n",myStrs,myStrs)
        fmt.Printf("Type %T is the same as type %T.\n",myStrs,strs)
        fmt.Println()
    }

    // 示例 2
    {
        type MyString string 
        str := "BCD"
        myStr1 := MyString(str)
        myStr2 := MyString("A"+str)
        _ = myStr2
        fmt.Printf("Type %T is different from type %T.\n",myStr1,str)

        strs := []string{"E","F","G"}

        var myStrs []MyString
        fmt.Printf("Type %T is different from type %T.\n",myStrs,strs)
        fmt.Println()
    }

    // 示例 3
    {
        type MyString1 = string 
        type MyString2 string
        str := "BCD"
        myStr1 := MyString1(str)
        myStr2 := MyString2(str)
        myStr1 = MyString1(myStr2)
        myStr2 = MyString2(myStr1)

        myStr1 = str
    }
}
