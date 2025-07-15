package atcoder.a7

fun main(){
    fun inputInts() = readLine()!!.split(" ").map(String::toInt)

    fun sol(X: List<Int>): Int {
        var s = X.map{it.toLong()}.sum()
        if(s % 3 != 0L){
            return -1
        }
        var m = (s/3).toInt()
        if(X.any{(m-it)%2 != 0}){
            return -1
        }
        return X.sumOf{ if(it > m) it-m else 0} / 2
    }
    val (T) = inputInts()
    repeat(T) {
        println(sol(inputInts()))
    }
}