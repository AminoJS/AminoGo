package main

import (
    "fmt"
    "github.com/AminoJS/AminoGo/aminogo"
    "os"
)

func main() {
    err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
    if err != nil {
        fmt.Println(err)
        return
    }
    
    yourFirstCommunity, err := aminogo.GetJoinedCommunities(aminogo.QueryOptions{
        Start: 0,
        Size:  1,
    })
    if err != nil {
        fmt.Println(err)
        return
    }
    
    blogs, err := aminogo.GetBlogFeed(yourFirstCommunity.CommunityList[0].NdcID, 0, 1)
    if err != nil {
        fmt.Println(err)
        return
    }
    blog := blogs.BlogList[0]
    fmt.Printf("Blog ID: %d\nContent:\n%s", blog.NdcID, blog.Content)
}
