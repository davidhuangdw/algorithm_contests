package atcoder.a5

fun main(){
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    fun sol(n: Int, p: List<Int>): Int{
        if(p.withIndex().all{(i, v) -> i+1 == v}) return 0
        var mx = 0
        for((i, v) in p.withIndex()){
            if(i+1 == v && mx == i){
                return 1
            }
            mx = maxOf(mx, v)
        }
        return if(p[0] == n && p.last() == 1) 3 else 2
    }

    val (T) = inputInts()
    repeat(T){
        val (N) = inputInts()
        val P = inputInts()
        println(sol(N, P))
    }
}
