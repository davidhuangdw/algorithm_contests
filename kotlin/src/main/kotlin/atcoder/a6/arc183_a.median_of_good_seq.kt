package atcoder.a6

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }

    val (N, K) = input_ints()
    if(N==1){
        println(List(K){1}.joinToString(" "))
        return
    }

    val p = N/2
    val q = N+1 - p
    val md = (N+1)/2

    val res = mutableListOf<Int>()
    if(md != p){
        res.addAll(List(K){md})
    }
    res.add(p)
    for(i in N downTo q){
        res.addAll(List(K){i})
    }
    res.addAll(List(K-1){p})
    for(i in p-1 downTo 1){
        res.addAll(List(K){i})
    }
    println(res.joinToString(" "))
}