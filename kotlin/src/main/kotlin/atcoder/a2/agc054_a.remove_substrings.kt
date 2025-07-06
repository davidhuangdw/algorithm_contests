package atcoder.a2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}
    fun input_longs() = split_input().mapNotNull{if(it != "") it.toLong() else null}

    val (n,) = input_ints()
    val S = readLine()!!
    val a = S[0]
    if (a != S[n-1]){
        println(1)
        return
    }
    if(n>= 4 && (1..n-3).any{i -> (S[i] != a && S[i+1] != a)}){
        println(2)
        return
    }
    println(-1)
}
