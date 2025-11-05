package atcoder.a11

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    val (N, M) = inputInts()
    val T = readLine()!!.toMutableList()
    val S = readLine()!!
    val cnt = IntArray(10)
    for(ch in S){
        cnt[ch-'0']+=1
    }
    var i = 0
    for(v in 9.downTo(1)){
        val ch = '0'+v
        while(i<N && cnt[v] > 0){
            if(T[i] < ch){
                T[i] = ch
                cnt[v] -= 1
            }
            i++
        }
    }
    if(!T.contains(S.last())){T[N-1] = S.last()} // bug: equal-value case: not-used but no need
    println(T.joinToString(""))
}
