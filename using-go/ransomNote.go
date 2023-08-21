package main

func canConstruct(ransomNote string, magazine string) bool {
    dictionary := make([]int, 26)
    for _, v := range magazine {
        dictionary[v-'a']++
    }
    for _, v := range ransomNote {
        dictionary[v-'a']--
        if dictionary[v-'a'] < 0 {
            return false
        }
    }
    return true
}