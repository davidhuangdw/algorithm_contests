package atcoder.a8

fun main() {
    readLine()!!
    val S = readLine()!!
    val T = readLine()!!
    val N = S.length
    var firstA = N
    var lastB = -1
    val aToB = mutableListOf<Int>()
    val bToA = mutableListOf<Int>()
    for ((i, p) in S.zip(T).withIndex()) {
        val (s, t) = p
        if (t == 'A') firstA = minOf(firstA, i)
        else lastB = maxOf(lastB, i)
        if (s != t) {
            if (s == 'A') aToB.add(i)
            else bToA.add(i)
        }
    }
//    println("${aToB} ${bToA} ${firstA} ${lastB}")
    if ((aToB.isNotEmpty() && aToB[0] < firstA) || (bToA.isNotEmpty() && bToA.last() > lastB)) {
        println(-1)
        return
    }

    var i = 0
    for (j in aToB) {
        if (i < bToA.size && bToA[i] < j) {
            i++
        }
    }
    println(aToB.size + bToA.size - i)
}