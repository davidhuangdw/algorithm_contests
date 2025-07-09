package atcoder.a3

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().mapNotNull{if(it != "") it.toInt() else null}

    readLine()
    val A = input_ints()
    var s = 0L
    var ss = 0L
    var mx = 0
    for((i, v) in A.withIndex()){
        s += v
        ss += s
        mx = maxOf(mx, v)
        println(1L*(i+1)*mx + ss)
    }
}