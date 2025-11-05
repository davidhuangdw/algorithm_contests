package atcoder.a12

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N,) = inputInts()
    val P = inputInts().toMutableList()
    val res = (0 until P.size step 2).flatMap { i ->
        val mx = maxOf(P[i], P[i+1])
        if(i+2 < P.size && P[i+2] > mx){
            P[i+1]=P[i+2].also{P[i+2]=P[i+1]}
            listOf(i+2)
        }else if(mx == P[i]){
            P[i]=P[i+1].also{P[i+1]=P[i]}
            listOf(i+1)
        }else{
            listOf()
        }
    }
    println(res.size)
    println(res.joinToString(" "))
}
