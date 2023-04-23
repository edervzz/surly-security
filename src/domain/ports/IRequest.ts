interface IRequest<TResponse> { }

interface IRequestHandler<TRequest extends TResponse, TResponse> {

    Handle(request: TRequest): TResponse

}