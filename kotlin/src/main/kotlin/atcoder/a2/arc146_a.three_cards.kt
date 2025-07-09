package atcoder.a2

fun main(){
    fun split_input(): List<String> = readLine()!!.trim().split(" ")
    fun input_ints() = split_input().map { it.toInt() }
    readLine()
//    val A = split_input().sortedByDescending { it.length }
//    val N = A.size
//    val l = A[2].length
//    val mx = mutableListOf<String>()
//    val k = (2 until N).find{ A[it].length < l } ?: N
//    var end = A.subList(2, k).toMutableList()
//    for(i in 0..1){
//        if(A[i].length == l){
//            end.add(A[i])
//        }else{
//            mx.add(A[i])
//        }
//    }
//    val rem = 3 - mx.size
//    println((end.sortedDescending().subList(0, rem) + mx).sortedDescending().joinToString("")) // bug: 1, 10 => not 101
    val th = input_ints().sortedDescending().take(3).map{it.toString()}
    var res = ""
    for(i in 0..2)
        for(j in 0..2)
            for(k in 0..2){
                if(setOf(i,j,k).size == 3){
                    res = maxOf(res, th[i]+th[j]+th[k])
                }
            }
    println(res)
}