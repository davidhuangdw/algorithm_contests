package atcoder.a8

fun main(){
    fun input_ints() = readLine()!!.split(" ").map { it.toInt() }

    val (N, ) = input_ints()
    var P = input_ints()
    var i = P.indexOf(1)
    if(P[(i+1)%N] == 2){
        println(minOf(i, N-i+2))
    }else{
        println(minOf(N-i, i+2))
    }
}