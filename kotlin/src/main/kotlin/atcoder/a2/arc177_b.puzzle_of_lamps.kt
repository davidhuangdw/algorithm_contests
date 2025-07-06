package atcoder.a2

fun main(){
    readLine()
    val S = readLine()!!
    val n = S.length
    var i = n-1
    val op = mutableListOf<String>()

    while(i >= 0){
        while(i>=0 && S[i] != '1'){
            i--
        }
        if(i < 0) break
        op.add("A".repeat(i+1))
        while(i>=0 && S[i] == '1'){
            i--
        }
        if(i < 0) break
        op.add("B".repeat(i+1))
    }
    val res = op.joinToString("")
    println(res.length)
    println(res)
}